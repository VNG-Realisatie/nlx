import styled from 'styled-components'

const Navigation = styled.div`
    display: flex;
    align-items: center;

    height: 64px;
    background-color: ${p => p.theme.color.white};

    box-shadow: 0 1px 4px rgba(0,0,0,.05);
`

export default Navigation
export { default as NavigationItem } from './NavigationItem'
