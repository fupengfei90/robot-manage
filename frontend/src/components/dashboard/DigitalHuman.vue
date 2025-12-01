<template>
  <el-card shadow="never" class="panel">
    <div class="panel__header">
      <h2>{{ t('dashboard.digitalHuman.title') }} <el-tag size="small" type="warning">TODO</el-tag></h2>
      <el-button text type="primary" @click="speak">{{ t('dashboard.digitalHuman.speak') }}</el-button>
    </div>
    <div ref="canvasRef" class="canvas"></div>
  </el-card>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref, watch } from 'vue'
import * as THREE from 'three'
import { useThemeStore } from '../../stores/theme'
import { useI18n } from '../../composables/useI18n'
import { useI18nStore } from '../../stores/i18n'

const canvasRef = ref<HTMLDivElement | null>(null)
const themeStore = useThemeStore()
const { t } = useI18n()
const i18nStore = useI18nStore()
let renderer: THREE.WebGLRenderer | null = null
let scene: THREE.Scene | null = null
let animationId = 0

const updateSceneBackground = () => {
  if (!scene) return
  const bgColor = getComputedStyle(document.documentElement)
    .getPropertyValue('--bg-secondary')
    .trim() || '#0f172a'
  scene.background = new THREE.Color(bgColor)
}

const initThree = () => {
  if (!canvasRef.value) return

  const width = canvasRef.value.clientWidth
  const height = 300

  scene = new THREE.Scene()
  updateSceneBackground()

  const camera = new THREE.PerspectiveCamera(45, width / height, 0.1, 1000)
  camera.position.z = 5

  renderer = new THREE.WebGLRenderer({ antialias: true })
  renderer.setSize(width, height)
  canvasRef.value.appendChild(renderer.domElement)

  const geometry = new THREE.BoxGeometry()
  const material = new THREE.MeshStandardMaterial({ color: '#3b82f6' })
  const cube = new THREE.Mesh(geometry, material)
  scene.add(cube)

  const light = new THREE.PointLight('#ffffff', 1)
  light.position.set(5, 5, 5)
  scene.add(light)

  const animate = () => {
    cube.rotation.x += 0.01
    cube.rotation.y += 0.01
    if (renderer && scene) {
      renderer.render(scene, camera)
    }
    animationId = requestAnimationFrame(animate)
  }

  animate()
}

const cleanup = () => {
  if (renderer) {
    renderer.dispose()
  }
  cancelAnimationFrame(animationId)
}

const speak = () => {
  const message = i18nStore.locale === 'zh'
    ? '智能助手机器人已就绪，持续监控生产系统。'
    : 'Digital assistant robot is ready, continuously monitoring production systems.'
  const utterance = new SpeechSynthesisUtterance(message)
  window.speechSynthesis.speak(utterance)
}

onMounted(initThree)
onUnmounted(cleanup)

watch(() => themeStore.mode, () => {
  updateSceneBackground()
})
</script>

<style scoped>
.panel {
  background: var(--bg-card);
  color: var(--text-primary);
  transition: background-color 0.3s ease, color 0.3s ease;
}
.canvas {
  width: 100%;
  height: 300px;
}

.panel__header :deep(.el-tag) {
  background-color: #fef3c7 !important;
  color: #92400e !important;
  border-color: #fbbf24 !important;
}
</style>
