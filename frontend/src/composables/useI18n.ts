import { useI18nStore } from '../stores/i18n'

export const useI18n = () => {
  const store = useI18nStore()
  return {
    t: store.t,
    locale: store.locale,
    setLocale: store.setLocale,
    toggle: store.toggle
  }
}

