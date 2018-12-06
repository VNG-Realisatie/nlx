import React, { PureComponent } from 'react'
import styled from 'styled-components'
import { Box } from '@rebass/grid'

const StyledContainer = styled(Box)`
    width: 100%;
    max-width: ${(p) => p.theme.containerWidth};
`

class Container extends PureComponent {
    render() {
        const { children } = this.props

        return <StyledContainer {...this.props}>{children}</StyledContainer>
    }
}

Container.propTypes = {}

Container.defaultProps = {
    px: '32px',
    mx: 'auto',
}

export default Container
