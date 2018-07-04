local k = import 'ksonnet/ksonnet.beta.3/k.libsonnet';

(import './alerts/kubermatic.libsonnet') +
(import './alerts/machine-controller.libsonnet') +
(import 'kubernetes-mixin/mixin.libsonnet') +
(import 'etcd-mixin/mixin.libsonnet') + {
  _config+:: {
    namespace: 'monitoring',

    // Selectors are inserted between {} in Prometheus queries.
    cadvisorSelector: 'job="cadvisor"',
    kubeletSelector: 'job="kubelet"',
    kubeStateMetricsSelector: 'job="kube-state-metrics"',
    nodeExporterSelector: 'app="node-exporter"',
    notKubeDnsSelector: 'job!="kube-dns"',
    kubeSchedulerSelector: 'job="kube-scheduler"',
    kubeControllerManagerSelector: 'job="kube-controller-manager"',
    kubeApiserverSelector: 'job="apiserver"',
    machineControllerSelector: 'job="machine-controller"',

    // We build alerts for the presence of all these jobs.
    jobs+:: {
      Cadvisor: $._config.cadvisorSelector,
      Kubelet: $._config.kubeletSelector,
      KubermaticAPI: 'job="pods",namespace="kubermatic",role="kubermatic-api"',
      KubermaticControllerManager: 'job="pods",namespace="kubermatic",role="controller-manager"',
      KubernetesApiserver: $._config.kubeApiserverSelector,
      KubeStateMetrics: $._config.kubeStateMetricsSelector,
      MachineController: $._config.machineControllerSelector,
      // KubernetesControllerManager: $._config.kubeControllerManagerSelector,
      // KubernetesScheduler: $._config.kubeSchedulerSelector,
    },

    runbookURLPattern: 'https://docs.kubermatic.io/monitoring/runbook/#alert-%s',

    prometheus+:: {
      name: 'kubermatic',
      rules: $.prometheusRules + $.prometheusAlerts,
    },
  },

  prometheus+:: {
    rules+:
      $._config.prometheus.rules,
  },
}
