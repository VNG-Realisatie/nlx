import styled from 'styled-components'

const Drawer = styled.div`
    display: flex;
    flex-direction: column;

    background-color: ${p => p.theme.color.white};
    box-shadow: 0 1px 4px rgba(0,0,0,.05);
`

export default Drawer