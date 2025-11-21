import { onMounted, onUnmounted, ref, watch } from 'vue'
import * as echarts from 'echarts'

export const useECharts = (options: echarts.EChartsCoreOption) => {
  const chartRef = ref<HTMLElement | null>(null)
  let instance: echarts.ECharts | null = null

  const init = () => {
    if (chartRef.value) {
      instance = echarts.init(chartRef.value)
      instance.setOption(options)
    }
  }

  const dispose = () => {
    instance?.dispose()
    instance = null
  }

  onMounted(init)
  onUnmounted(dispose)

  watch(
    () => options,
    (val) => {
      if (instance && val) {
        instance.setOption(val)
        instance.resize()
      }
    },
    { deep: true }
  )

  return { chartRef, instance }
}
