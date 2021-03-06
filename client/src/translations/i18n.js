import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'

import { TRANSLATIONS_ES } from './es'

i18n.use(initReactI18next).init({
  resources: {
    es: {
      translation: TRANSLATIONS_ES,
    },
  },
  lgn: 'es',
  fallbackLng: 'es',
})
