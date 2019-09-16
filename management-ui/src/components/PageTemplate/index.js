// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { Container, Main, Navbar } from './index.styles'

const PageTemplate = ({ children }) => (
    <Container>
        <Navbar />
        <Main role="main">
            <Main.Container>{children}</Main.Container>
        </Main>
    </Container>
)

export default PageTemplate
