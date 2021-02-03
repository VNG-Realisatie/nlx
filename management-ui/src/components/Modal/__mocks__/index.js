// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
const modalExports = jest.requireActual('../index')
export default modalExports.default
export { modalExports }

// This mock rerenders because of useEffect
// So after triggering a `close` you need to wrap an assertion in waitFor
//
// fireEvent.click(getByText('close modal button'))
// await waitFor(() => expect(clostHandler).toHaveBeenCalled())
//
jest.mock('react-transition-group', () => {
  const React = require('react')
  const { useRef, useEffect } = React

  const FakeTransition = jest.fn((props) => {
    const inProp = useRef(props.in)

    useEffect(() => {
      if (props.in === inProp.current) return

      // Modal only uses onExited
      if (!props.in && props.onExited) {
        props.onExited()
      }

      inProp.current = props.in
    }, [props])

    return props.children
  })

  const FakeCSSTransition = jest.fn((props) => (
    <FakeTransition {...props}>{props.children}</FakeTransition>
  ))

  return { CSSTransition: FakeCSSTransition, Transition: FakeTransition }
})
