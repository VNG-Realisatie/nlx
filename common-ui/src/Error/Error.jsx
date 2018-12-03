import styled from 'styled-components'

const Error = styled.div`
    margin-top: 0.5rem;
    font-size: ${(p) => p.theme.font.size.small};
    line-height: ${(p) => p.theme.font.lineHeight.small};
    color: ${(p) => p.theme.color.alert};
`

export default Error
