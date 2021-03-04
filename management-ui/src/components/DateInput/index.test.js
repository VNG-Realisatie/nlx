// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { dateToIsoFormat, isoDateSchema } from './index'

test('dateToIsoFormat', () => {
  expect(dateToIsoFormat(new Date('2020-2-2'))).toEqual('2020-02-02')
})

test('isoDateSchema', async () => {
  const schema = isoDateSchema('My error')

  expect(await schema.isValid('2020-02-02')).toBe(true)
  expect(await schema.isValid('2020-02-2')).toBe(false)
  expect(await schema.isValid('02-02-2020')).toBe(false)

  expect(await schema.validate('asdf').catch((err) => err.message)).toEqual(
    'My error',
  )

  expect(await schema.isValid('')).toBe(true)
  expect(await schema.required().isValid('')).toBe(false)
})
