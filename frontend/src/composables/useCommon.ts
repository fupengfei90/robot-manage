import { useI18n } from './useI18n'
import { useThemeStore } from '../stores/theme'

export const useCommon = () => {
  const { t, locale } = useI18n()
  const themeStore = useThemeStore()

  return {
    t,
    locale,
    themeStore
  }
}
