package wasmtype

import (
    "context"
    "github.com/crossplane/crossplane-runtime/pkg/controller"
    "github.com/crossplane/crossplane-runtime/pkg/logging"
    "github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
    "sigs.k8s.io/controller-runtime/pkg/manager"

    v1alpha1 "github.com/lucazini03/my-provider/apis/wasmgroup/v1alpha1"
    corev1 "k8s.io/api/core/v1"
    v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Setup aggiunge un controller per il tipo Wasm.
func Setup(mgr manager.Manager, l logging.Logger) error {
    name := managed.ControllerName(v1alpha1.WASMTypeGroupKind)

    r := managed.NewReconciler(mgr,
        managed.NewAPIManaged(mgr.GetClient(), &v1alpha1.WASMType{}),
        managed.WithExternalConnecter(&connector{
            kube:    mgr.GetClient(),
            newFn:   NewExternal,
        }),
    )

    return ctrl.NewControllerManagedBy(mgr).
        Named(name).
        For(&v1alpha1.WASMType{}).
        Complete(r)
}

type connector struct {
    kube client.Client
    newFn func() *external
}

func (c *connector) Connect(ctx context.Context, mg managed.Managed) (managed.ExternalClient, error) {
    // Connessione a Kubernetes per creare il Pod con runtimeClassName wasmedge
    return c.newFn(), nil
}

type external struct {}

func NewExternal() *external {
    return &external{}
}

func (e *external) Observe(ctx context.Context, mg managed.Managed) (managed.ExternalObservation, error) {
    // Logica di osservazione del Pod Wasm
    return managed.ExternalObservation{}, nil
}

func (e *external) Create(ctx context.Context, mg managed.Managed) (managed.ExternalCreation, error) {
    cr, ok := mg.(*v1alpha1.WASMType)
    if !ok {
        return managed.ExternalCreation{}, nil
    }

    // Creazione del Pod Wasm
    pod := &corev1.Pod{
        ObjectMeta: v1.ObjectMeta{
            Name:      cr.GetName(),
            Namespace: cr.GetNamespace(),
        },
        Spec: corev1.PodSpec{
            Containers: []corev1.Container{
                {
                    Name:  "wasm-module",
                    Image: cr.Spec.Image,
                },
            },
            RuntimeClassName: &cr.Spec.RuntimeClassName,
        },
    }

    if err := controllerutil.SetControllerReference(cr, pod, client.Scheme()); err != nil {
        return managed.ExternalCreation{}, err
    }

    err := e.kube.Create(ctx, pod)
    if err != nil {
        return managed.ExternalCreation{}, err
    }

    return managed.ExternalCreation{}, nil
}
