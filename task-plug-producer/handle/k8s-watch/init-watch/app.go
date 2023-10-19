package init_watch

import (
	v1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type App struct {
	service     []*corev1.Service
	statefulSet []*v1.StatefulSet
	deployment  []*v1.Deployment
	pod         []*corev1.Pod
	cm          []*corev1.ConfigMap
	pvc         []*corev1.PersistentVolumeClaim
	job         []*batchv1.Job
}

func InitCacheApp() *App {
	return &App{}
}

func (a *App) SetDeployment(d *v1.Deployment) {
	if len(a.deployment) > 0 {
		for i, deploy := range a.deployment {
			if deploy.GetName() == d.GetName() {
				a.deployment[i] = d
				return
			}
		}
	}
	a.deployment = append(a.deployment, d)
}

func (a *App) DeleteDeployment(d *v1.Deployment) {
	for i, old := range a.deployment {
		if old.GetName() == d.GetName() {
			a.deployment = append(a.deployment[0:i], a.deployment[i+1:]...)
			return
		}
	}
}

func (a *App) GetDeployment() []*v1.Deployment {
	if a.deployment != nil {
		return a.deployment
	}
	return []*v1.Deployment{}
}

func (a *App) DeleteStatefulSet(d *v1.StatefulSet) {
	for i, old := range a.statefulSet {
		if old.GetName() == d.GetName() {
			a.statefulSet = append(a.statefulSet[0:i], a.statefulSet[i+1:]...)
			return
		}
	}
}

func (a *App) SetStatefulSet(d *v1.StatefulSet) {
	if len(a.statefulSet) > 0 {
		for i, sts := range a.statefulSet {
			if sts.GetName() == d.GetName() {
				a.statefulSet[i] = d
				return
			}
		}
	}
	a.statefulSet = append(a.statefulSet, d)
}

func (a *App) GetStatefulSet() []*v1.StatefulSet {
	if a.statefulSet != nil {
		return a.statefulSet
	}
	return []*v1.StatefulSet{}
}

func (a *App) DeleteJob(d *batchv1.Job) {
	for i, old := range a.job {
		if old.GetName() == d.GetName() {
			a.job = append(a.job[0:i], a.job[i+1:]...)
			return
		}
	}
}

func (a *App) SetJob(d *batchv1.Job) {
	if len(a.job) > 0 {
		for i, j := range a.job {
			if j.GetName() == d.GetName() {
				a.job[i] = d
				return
			}
		}
	}
	a.job = append(a.job, d)
}

func (a *App) GetJob() []*batchv1.Job {
	if a.job != nil {
		return a.job
	}
	return []*batchv1.Job{}
}

func (a *App) DeleteService(d *corev1.Service) {
	for i, old := range a.service {
		if old.GetName() == d.GetName() {
			a.service = append(a.service[0:i], a.service[i+1:]...)
			return
		}
	}
}

func (a *App) SetService(d *corev1.Service) {
	if len(a.service) > 0 {
		for i, svc := range a.service {
			if svc.GetName() == d.GetName() {
				a.service[i] = d
				return
			}
		}
	}
	a.service = append(a.service, d)
}

func (a *App) GetService() []*corev1.Service {
	if a.statefulSet != nil {
		return a.service
	}
	return []*corev1.Service{}
}

func (a *App) DeletePod(d *corev1.Pod) {
	for i, old := range a.pod {
		if old.GetName() == d.GetName() {
			a.pod = append(a.pod[0:i], a.pod[i+1:]...)
			return
		}
	}
}

func (a *App) SetPod(d *corev1.Pod) {
	if len(a.pod) > 0 {
		for i, pod := range a.pod {
			if pod.GetName() == d.GetName() {
				a.pod[i] = d
				return
			}
		}
	}
	a.pod = append(a.pod, d)
}

func (a *App) GetPod() []*corev1.Pod {
	if a.pod != nil {
		return a.pod
	}
	return []*corev1.Pod{}
}

func (a *App) DeleteConfigMap(d *corev1.ConfigMap) {
	for i, old := range a.cm {
		if old.GetName() == d.GetName() {
			a.cm = append(a.cm[0:i], a.cm[i+1:]...)
			return
		}
	}
}

func (a *App) SetConfigMap(d *corev1.ConfigMap) {
	if len(a.cm) > 0 {
		for i, cm := range a.cm {
			if cm.GetName() == d.GetName() {
				a.cm[i] = d
				return
			}
		}
	}
	a.cm = append(a.cm, d)
}

func (a *App) GetConfigMap() []*corev1.ConfigMap {
	if a.cm != nil {
		return a.cm
	}
	return []*corev1.ConfigMap{}
}

func (a *App) DeletePVC(d *corev1.PersistentVolumeClaim) {
	for i, old := range a.pvc {
		if old.GetName() == d.GetName() {
			a.pvc = append(a.pvc[0:i], a.pvc[i+1:]...)
			return
		}
	}
}

func (a *App) SetPVC(d *corev1.PersistentVolumeClaim) {
	if len(a.pvc) > 0 {
		for i, pvc := range a.pvc {
			if pvc.GetName() == d.GetName() {
				a.pvc[i] = d
				return
			}
		}
	}
	a.pvc = append(a.pvc, d)
}

func (a *App) GetPVC() []*corev1.PersistentVolumeClaim {
	if a.pvc != nil {
		return a.pvc
	}
	return []*corev1.PersistentVolumeClaim{}
}
