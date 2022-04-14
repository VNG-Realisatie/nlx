// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { render } from '@testing-library/react'

// code from https://github.com/testing-library/react-hooks-testing-library/issues/654#issuecomment-1097276573
// this code can be removed once https://github.com/testing-library/react-testing-library/pull/991
// has been resolved
function renderHook(renderCallback, options = {}) {
  const { initialProps, wrapper } = options
  const result = React.createRef()

  function TestComponent({ renderCallbackProps }) {
    const pendingResult = renderCallback(renderCallbackProps)

    React.useEffect(() => {
      result.current = pendingResult
    })

    return null
  }

  const { rerender: baseRerender, unmount } = render(
    <TestComponent renderCallbackProps={initialProps} />,
    { wrapper },
  )

  function rerender(rerenderCallbackProps) {
    return baseRerender(
      <TestComponent renderCallbackProps={rerenderCallbackProps} />,
    )
  }

  return { result, rerender, unmount }
}

export default renderHook
