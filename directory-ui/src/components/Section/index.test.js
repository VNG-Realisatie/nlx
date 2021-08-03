// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { getColor } from './index'

const props = {
  theme: {
    tokens: {
      colorBackground: 'hotpink',
    },
    colorAlternateSection: 'lime',
  },
}

test('getColor returns expected color', () => {
  expect(getColor({ ...props, alternate: false })).toEqual('hotpink')
  expect(getColor({ ...props, alternate: false }, true)).toEqual('lime')
  expect(getColor({ ...props, alternate: true })).toEqual('lime')
  expect(getColor({ ...props, alternate: true }, true)).toEqual('hotpink')
})
