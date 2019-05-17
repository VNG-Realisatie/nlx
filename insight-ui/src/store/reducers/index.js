// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'

import loader from './loader'
import i18n from './i18n'
import organizations from './organizations'
import organization from './organization'

export default combineReducers({
    loader,
    i18n,
    organizations,
    organization,
})
