// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { fireEvent, waitFor } from '@testing-library/react'

export function clickConfirmButton(button) {
  jest.useFakeTimers()
  fireEvent.click(button)
  jest.runAllTimers()
  jest.useRealTimers()
}

export function clickConfirmButtonAndAssert(button, assertion) {
  clickConfirmButton(button)
  return waitFor(assertion)
}
