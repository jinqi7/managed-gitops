package appstudioredhatcom

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/codeready-toolchain/toolchain-common/pkg/condition"

	codereadytoolchainv1alpha1 "github.com/codeready-toolchain/api/api/v1alpha1"
)

// SpaceRequestReadyPredicate returns a predicate which filters out
// only SpaceRequests whose Ready Status are set to true
func SpaceRequestReadyPredicate() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(createEvent event.CreateEvent) bool {
			return false
		},
		DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
			return false
		},
		GenericFunc: func(genericEvent event.GenericEvent) bool {
			return false
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return hasSpaceRequestChangedToReady(e.ObjectOld, e.ObjectNew)
		},
	}
}

//IsSpaceRequestReady checks if SpaceRequest condition is in Ready status.
func IsSpaceRequestReady(spacerequest *codereadytoolchainv1alpha1.SpaceRequest) bool {
	return condition.IsTrue(spacerequest.Status.Conditions, codereadytoolchainv1alpha1.ConditionReady)
}

// hasSpaceRequestChangedToReady returns a boolean indicating whether the SpaceRequest becomes Ready.
// If the objects passed to this function are not SpaceRequest, the function will return false.
func hasSpaceRequestChangedToReady(objectOld, objectNew client.Object) bool {
	if newSpaceRequest, ok := objectNew.(*codereadytoolchainv1alpha1.SpaceRequest); ok {
		return IsSpaceRequestReady(newSpaceRequest) && (len(newSpaceRequest.Status.NamespaceAccess) >= 1)
	}
	return false
}
