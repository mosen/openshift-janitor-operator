package sweep

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"time"

	comv1alpha1 "github.com/mosen/openshift-janitor-operator/pkg/apis/janitor/v1alpha1"
	projectv1 "github.com/openshift/api/project/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_sweep")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Sweep Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	projectv1.Install(mgr.GetScheme())
	return &ReconcileSweep{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("sweep-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Sweep
	err = c.Watch(&source.Kind{Type: &comv1alpha1.Sweep{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Sweep
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &comv1alpha1.Sweep{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileSweep implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileSweep{}

// ReconcileSweep reconciles a Sweep object
type ReconcileSweep struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Sweep object and makes changes based on the state read
// and what is in the Sweep.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSweep) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Sweep")

	// Fetch the Sweep instance
	instance := &comv1alpha1.Sweep{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	_, err = r.execSweep(instance)
	if err != nil {
		reqLogger.Error(err, "Unable to execute Sweep")
		return reconcile.Result{}, err
	}
	// Set Sweep instance as the owner and controller
	//if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
	//	return reconcile.Result{}, err
	//}

	//// Check if this Pod already exists
	//found := &corev1.Pod{}
	//err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	//if err != nil && errors.IsNotFound(err) {
	//	reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
	//	err = r.client.Create(context.TODO(), pod)
	//	if err != nil {
	//		return reconcile.Result{}, err
	//	}
	//
	//	// Pod created successfully - don't requeue
	//	return reconcile.Result{}, nil
	//} else if err != nil {
	//	return reconcile.Result{}, err
	//}
	//
	//// Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// Exec sweep executes a sweep over the Projects in the current cluster
func (r *ReconcileSweep) execSweep(cr *comv1alpha1.Sweep) ([]string, error) {
	now := metav1.NewTime(time.Now())
	//cr.Status.Active = true
	// cr.Status.Started = &now
	//r.client.Update(context.TODO(), cr)

	oldestTimestamp := metav1.NewTime(now.AddDate(0, 0, -1*cr.Spec.MaximumAgeDays))

	projects := &projectv1.ProjectList{}
	opts := []client.ListOption{
		client.InNamespace(""),
	}
	if err := r.client.List(context.TODO(), projects, opts...); err != nil {
		return nil, err
	}

	if len(projects.Items) == 0 {
		return nil, nil
	}

	sweepLogger := log.WithValues()

OuterLoop:
	for _, project := range projects.Items {
		if strings.HasPrefix(project.GetName(), "openshift") || strings.HasPrefix(project.GetName(), "kube-") {
			sweepLogger.Info("skipping system namespace/project", "metav1.Name", project.GetName())
			continue
		}

		for _, ignoreProject := range cr.Spec.IgnoreProjects {
			if ignoreProject == project.GetName() {
				sweepLogger.Info("skipping ignored project", "metav1.Name", project.GetName())
				break OuterLoop
			}
		}

		creationTs := project.GetCreationTimestamp()
		sweepLogger.Info("processing project", "metav1.Name", project.GetName(), "metav1.CreationTimestamp", creationTs)
		if creationTs.Before(&oldestTimestamp) {
			difference := creationTs.Sub(oldestTimestamp.Time)
			sweepLogger.Info("project is older than MaximumAgeDays", "metav1.Name", project.GetName(), "ageDays", difference.Hours()/24)
		}
	}

	return nil, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
//func newPodForCR(cr *comv1alpha1.Sweep) *corev1.Pod {
//	labels := map[string]string{
//		"app": cr.Name,
//	}
//	return &corev1.Pod{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      cr.Name + "-pod",
//			Namespace: cr.Namespace,
//			Labels:    labels,
//		},
//		Spec: corev1.PodSpec{
//			Containers: []corev1.Container{
//				{
//					Name:    "busybox",
//					Image:   "busybox",
//					Command: []string{"sleep", "3600"},
//				},
//			},
//		},
//	}
//}
