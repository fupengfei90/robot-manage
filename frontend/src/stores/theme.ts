import { defineStore } from 'pinia'

export type ThemeMode = 'dark' | 'light'

export const useThemeStore = defineStore('theme', {
  state: (): { mode: ThemeMode } => ({
    mode: (localStorage.getItem('theme_mode') as ThemeMode) || 'dark'
  }),
  actions: {
    toggle() {
      this.mode = this.mode === 'dark' ? 'light' : 'dark'
      localStorage.setItem('theme_mode', this.mode)
      this.applyTheme()
    },
    setMode(mode: ThemeMode) {
      this.mode = mode
      localStorage.setItem('theme_mode', mode)
      this.applyTheme()
    },
    applyTheme() {
      const root = document.documentElement
      if (this.mode === 'light') {
        root.classList.add('light-theme')
        root.classList.remove('dark-theme')
      } else {
        root.classList.add('dark-theme')
        root.classList.remove('light-theme')
      }
    }
  }
})

