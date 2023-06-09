/*
Copyright 2023 stackzoo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	examplesv1alpha1 "github.com/stackzoo/operator-blueprint/api/v1alpha1"
)

// PodBusterReconciler reconciles a PodBuster object
type PodBusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// RBAC permissions to monitor podbusters
//+kubebuilder:rbac:groups=examples.stackzoo.io,resources=podbusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=examples.stackzoo.io,resources=podbusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=examples.stackzoo.io,resources=podbusters/finalizers,verbs=update

// RBAC permissions to monitor pods
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;update;patch;delete

func (r *PodBusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("PodBuster - enter reconcile", "req", req)

	PodBuster := &examplesv1alpha1.PodBuster{}
	if err := r.Get(ctx, req.NamespacedName, PodBuster); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	l.Info("PodBuster", "PodBuster", PodBuster)

	// Get pod list
	pods := &corev1.PodList{}
	if err := r.List(ctx, pods, client.InNamespace(PodBuster.Spec.Namespace)); err != nil {
		l.Error(err, "Failed to list pods")
		return ctrl.Result{}, err
	}

	// Loop on pods and delete the ones that are in the namespace specified in PodBuster CR
	for _, pod := range pods.Items {
		// Remove the pod
		l.Info("PodBuster", "Deleting pod", pod.Name)
		if err := r.Delete(ctx, &pod); err != nil {
			l.Error(err, "Failed to delete pod", "pod", pod.Name)
			return ctrl.Result{}, err
		}
	}
	// Change PodBuster status
	PodBuster.Status.Ok = true
	if err := r.Status().Update(ctx, PodBuster); err != nil {
		l.Error(err, "unable to update PodBuster's status", "status", true)
		return ctrl.Result{}, err
	}
	l.Info("PodBuster's status updated", "status", true)
	l.Info("PodBuster custom resource reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodBusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplesv1alpha1.PodBuster{}).
		Owns(&examplesv1alpha1.PodBuster{}).
		Complete(r)
}
