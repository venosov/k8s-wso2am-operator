package patternX

import (
	apimv1alpha1 "github.com/wso2-incubator/wso2am-k8s-operator/pkg/apis/apim/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"strings"
)


func getApimXVolumes(apimanager *apimv1alpha1.APIManager, r apimv1alpha1.Profile, x *configvalues) ([]corev1.VolumeMount, []corev1.Volume) {

	var amXvolumemounts []corev1.VolumeMount
	var amXvolume []corev1.Volume
	defaultdeployConf :=  r.Deployment.Configmaps.DeploymentConfigmap

	// adding default deploymentConfigmap
	amXvolumemounts= append(amXvolumemounts,corev1.VolumeMount{
		Name: "wso2am-px-apim-conf",
		MountPath: "/home/wso2carbon/wso2-config-volume/repository/conf/deployment.toml",
		SubPath:"deployment.toml",
	})

	amXvolume = append(amXvolume,corev1.Volume{
		Name: "wso2am-px-apim-conf",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: defaultdeployConf,
				},
			},
		},
	})

	apimpvc := r.Deployment.PersistentVolumeClaim

	if apimpvc.SynapseConfigs != "" {

		defaultsynapseconf := "wso2am-p1-am-synapse-configs"

		var carbonHome = "/home/wso2carbon/"
		var synapseDir = "/repository/deployment/server/synapse-configs"

		s := []string{carbonHome, x.APIMVersion, synapseDir};
		var synapseLocation = strings.Join(s,"")

		// adding default synapseConfigs pvc
		amXvolumemounts = append(amXvolumemounts,corev1.VolumeMount{
			Name:      "wso2am-px-synapse-conf",
			MountPath: synapseLocation,

		})

		amXvolume = append(amXvolume,corev1.Volume{
			Name: "wso2am-px-synapse-conf",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName:defaultsynapseconf,
				},
			},
		})

	}

	if apimpvc.ExecutionPlans != "" {

		defaultexecutionconf :=	"wso2am-p1-am-execution-plans"

		var carbonHome = "/home/wso2carbon/"
		var executionPlansDir = "/repository/deployment/server/executionplans"

		ss := []string{carbonHome, x.APIMVersion, executionPlansDir};
		var executionPlansLocation = strings.Join(ss,"")

		// adding default executionPlans pvc
		amXvolumemounts=append(amXvolumemounts,corev1.VolumeMount{
			Name:        "wso2am-px-execution-conf",
			MountPath: 	 executionPlansLocation,
		})

		amXvolume =append(amXvolume,corev1.Volume{
			Name: "wso2am-px-execution-conf",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName:defaultexecutionconf,
				},
			},
		})

	}

	return amXvolumemounts, amXvolume
}

func getDashboardXVolumes(apimanager *apimv1alpha1.APIManager, r apimv1alpha1.Profile) ([]corev1.VolumeMount, []corev1.Volume) {

	var dashxvolumemounts []corev1.VolumeMount
	var dashxvolume []corev1.Volume

	dashpvc := r.Deployment.PersistentVolumeClaim
	dashVolDefined := dashpvc.SynapseConfigs != "" && dashpvc.ExecutionPlans != ""

	if dashVolDefined {

		defaultdashconf :=  r.Deployment.Configmaps.DeploymentConfigmap

		dashxvolumemounts=append(dashxvolumemounts,corev1.VolumeMount{
		Name: defaultdashconf,
		MountPath: "/home/wso2carbon/wso2-config-volume/conf/dashboard/deployment.yaml",
		SubPath:"deployment.yaml",


		})

		dashxvolume =append(dashxvolume,corev1.Volume{
			Name: defaultdashconf,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: defaultdashconf,
					},
				},
			},
		})
	}

	return dashxvolumemounts, dashxvolume

}


func getWorkerXVolumes(apimanager *apimv1alpha1.APIManager, r apimv1alpha1.Profile) ([]corev1.VolumeMount, []corev1.Volume) {

	var workerxvolumemounts []corev1.VolumeMount
	var workerxvolume []corev1.Volume

	workerpvc := r.Deployment.PersistentVolumeClaim
	workerVolDefined := workerpvc.SynapseConfigs != "" && workerpvc.ExecutionPlans != ""

	if workerVolDefined {

		defaultdeployConf :=  r.Deployment.Configmaps.DeploymentConfigmap		
		// adding default deploymentConfigmap
		workerxvolumemounts=append(workerxvolumemounts,corev1.VolumeMount{
			Name: defaultdeployConf,
			MountPath: "/home/wso2carbon/wso2-config-volume/conf/worker",
			//SubPath:"deployment.yaml",

		})

		workerxvolume =append(workerxvolume,corev1.Volume{
			Name: defaultdeployConf,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: defaultdeployConf,
					},
				},
			},
		})
	}

	return workerxvolumemounts, workerxvolume

}

