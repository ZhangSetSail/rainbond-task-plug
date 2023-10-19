package init_watch

import (
	"github.com/goodrain/rainbond-task-plug/pkg"
	"k8s.io/client-go/informers"
	appsv1 "k8s.io/client-go/listers/apps/v1"
	batchv1 "k8s.io/client-go/listers/batch/v1"
	corev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

func CreateResourceWatch() ManagerWatch {
	clientSet := pkg.GetClientSet()
	mw := &managerWatch{
		informers: &Informer{},
		listers:   &Lister{},
	}
	infFactory := informers.NewSharedInformerFactoryWithOptions(clientSet, 0, informers.WithNamespace(""))

	mw.listers.Service = infFactory.Core().V1().Services().Lister()
	mw.informers.Service = infFactory.Core().V1().Services().Informer()

	mw.listers.StatefulSet = infFactory.Apps().V1().StatefulSets().Lister()
	mw.informers.StatefulSet = infFactory.Apps().V1().StatefulSets().Informer()

	mw.informers.Deployment = infFactory.Apps().V1().Deployments().Informer()
	mw.listers.Deployment = infFactory.Apps().V1().Deployments().Lister()

	mw.informers.Pod = infFactory.Core().V1().Pods().Informer()
	mw.listers.Pod = infFactory.Core().V1().Pods().Lister()

	mw.informers.ConfigMap = infFactory.Core().V1().ConfigMaps().Informer()
	mw.listers.ConfigMap = infFactory.Core().V1().ConfigMaps().Lister()

	mw.informers.PVC = infFactory.Core().V1().PersistentVolumeClaims().Informer()
	mw.listers.PVC = infFactory.Core().V1().PersistentVolumeClaims().Lister()

	mw.informers.Job = infFactory.Batch().V1().Jobs().Informer()
	mw.listers.Job = infFactory.Batch().V1().Jobs().Lister()

	mw.informers.Service.AddEventHandler(mw)
	mw.informers.StatefulSet.AddEventHandler(mw)
	mw.informers.Deployment.AddEventHandler(mw)
	mw.informers.Pod.AddEventHandler(mw)
	mw.informers.ConfigMap.AddEventHandler(mw)
	mw.informers.PVC.AddEventHandler(mw)
	mw.informers.Job.AddEventHandler(mw)

	return mw
}

type Informer struct {
	Service     cache.SharedIndexInformer
	StatefulSet cache.SharedIndexInformer
	Deployment  cache.SharedIndexInformer
	Pod         cache.SharedIndexInformer
	ConfigMap   cache.SharedIndexInformer
	PVC         cache.SharedIndexInformer
	Job         cache.SharedIndexInformer
}

func (i *Informer) Start(stop chan struct{}) {
	go i.Service.Run(stop)
	go i.StatefulSet.Run(stop)
	go i.Deployment.Run(stop)
	go i.Pod.Run(stop)
	go i.ConfigMap.Run(stop)
	go i.PVC.Run(stop)
	go i.Job.Run(stop)
}

func (i *Informer) Ready() bool {
	if i.Service.HasSynced() && i.StatefulSet.HasSynced() && i.Deployment.HasSynced() && i.Pod.HasSynced() &&
		i.ConfigMap.HasSynced() && i.PVC.HasSynced() && i.Job.HasSynced() {
		return true
	}
	return false
}

type Lister struct {
	Service     corev1.ServiceLister
	StatefulSet appsv1.StatefulSetLister
	Deployment  appsv1.DeploymentLister
	Pod         corev1.PodLister
	ConfigMap   corev1.ConfigMapLister
	PVC         corev1.PersistentVolumeClaimLister
	Job         batchv1.JobLister
}
