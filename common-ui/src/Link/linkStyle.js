import { css } from 'styled-components'

export const linkStyle = css`
    display: inline-block;
    color: ${(p) => p.theme.color.primary.main};
    text-decoration: none;

    ${(p) =>
        p.underline &&
        css`
            border-bottom: 1px solid transparent;

            &:hover,
            &:focus {
                border-color: ${(p) => p.theme.color.primary.light};
            }
        `};
`
