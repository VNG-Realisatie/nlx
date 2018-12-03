import styled from 'styled-components'

const Small = styled.div`
    font-size: ${p => p.theme.font.size.small};
    line-height: ${p => p.theme.font.lineHeight.small};
    color: ${p => p.theme.color.grey[60]};
`

export default Small