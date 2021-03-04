// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { string } from 'yup'

const isoDateSchema = (message = 'Invalid date') =>
  string().test('isoDate', message, function (value = '') {
    if (value === '') return true
    if (!value.match(/\d{4}-\d{2}-\d{2}/)) return false

    const d = new Date(value)
    return d && d.getTime && !isNaN(d.getTime())
  })

const padZero = (n) => `0${n}`.slice(-2)

const dateToIsoFormat = (d) =>
  `${d.getFullYear()}-${padZero(d.getMonth() + 1)}-${padZero(d.getDate())}`

export { isoDateSchema, dateToIsoFormat }
