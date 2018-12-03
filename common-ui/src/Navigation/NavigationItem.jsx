import styled from 'styled-components'

const NavigationItem = styled.a`
    display: inline-flex;
    align-items: center;

    height: 64px;
    padding: 0 1rem;

    color: ${p => p.theme.color.grey[60]};

    text-decoration: none;
    text-shadow: none;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    user-select: none;

    transition: ${p => `background-color ${p.theme.transition.fast}`};

    &:hover,
    &:focus {
        background-color: ${p => p.theme.color.hover};
    }

    &:active {
        background-color: ${p => p.theme.color.active};
    }

    &[aria-current] {
        color: ${p => p.theme.color.primary.main};
    }
`

export default NavigationItem