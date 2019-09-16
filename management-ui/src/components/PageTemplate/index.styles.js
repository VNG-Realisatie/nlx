// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import styled from 'styled-components'
import UnstyledSidebar from '../Sidebar'

export const Container = styled.div`
    display: flex;
    height: 100vh;
`

const StyledMain = styled.main`
    flex: 1;
    overflow: auto;
    padding: 36px 24px 24px 24px;
    background-color: ${(p) => p.theme.color.body};
`

StyledMain.Container = styled.div`
    width: 100%;
    max-width: 860px;
    margin: 0 auto;
`

export const Main = StyledMain

export const Navbar = styled(UnstyledSidebar)`
    width: 250px;
    box-shadow: 1px 0 0 #e6eaf5;
`
