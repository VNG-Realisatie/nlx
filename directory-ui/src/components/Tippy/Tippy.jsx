import React, { Fragment } from 'react'
import TippyPlugin from '@tippy.js/react'

import { createGlobalStyle } from 'styled-components'
import tippyStyle from './tippyStyle'

const TippyStyle = createGlobalStyle`
    ${tippyStyle};
`

const Tippy = ({ children, ...rest }) => (
    <Fragment>
        <TippyStyle />
        <TippyPlugin {...rest}>{children}</TippyPlugin>
    </Fragment>
)

Tippy.defaultProps = {
    ignoreAttributes: true,
    duration: 200,
}

export default Tippy
