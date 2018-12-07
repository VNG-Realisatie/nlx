import media, { down, up, only } from '../mediaQueries'

jest.mock('../../themeConstants', () => ({
    breakpoints: {
        breakpoints: {
            sm: 0,
            md: 768,
            lg: 1024,
        },
        unit: 'px',
        step: 5, // Note: this is divided by 100 in the code
    },
}))

jest.mock('styled-components', () => ({
    // For the template literal functions, css`` is called twice, we receive a nested array like:
    // [[<handles>], <media query>, [[<passed css>]]]
    css: jest.fn((...args) => args),
}))

describe('Given theme/helpers/mediaQueries', () => {
    it('Should return an object containing expected functions', () => {
        const expected = [
            'smDown',
            'sm',
            'smUp',
            'mdDown',
            'md',
            'mdUp',
            'lgDown',
            'lg',
            'lgUp',
        ]
        const keys = Object.keys(media)

        expect(keys).toEqual(expected)
    })

    describe('The media query string functions', () => {
        it("Should return expected query for up('md')", () => {
            expect(up('md')).toEqual('@media (min-width: 768px)')
        })

        it("Should return expected query for down('md')", () => {
            // Note: this uses `${step} / 100`
            expect(down('md')).toEqual('@media (max-width: 1023.95px)')
        })

        it("Should return expected query for only('md')", () => {
            // Note: this uses `${step} / 100`
            expect(only('md')).toEqual(
                '@media (min-width: 768px) and (max-width: 1023.95px)',
            )
        })

        it("Should return expected query for down('lg')", () => {
            expect(down('lg')).toEqual('@media (min-width: 0px)')
        })

        it("Should return expected query for only('lg')", () => {
            expect(only('lg')).toEqual('@media (min-width: 1024px)')
        })

        it("Should return expected query for down('sm')", () => {
            expect(down('sm')).toEqual('@media (max-width: 767.95px)')
        })
    })

    describe('The template literals', () => {
        it('Should return the expected media query for `lg`', () => {
            const result = media.lg`color: red;`
            const query = result[1]

            expect(query).toEqual('@media (min-width: 1024px)')
        })

        it('Should return the expected media query for `mdDown`', () => {
            const result = media.mdDown`color: red;`
            const query = result[1]

            expect(query).toEqual('@media (max-width: 1023.95px)')
        })

        it('Should return the expected media query for `md`', () => {
            const result = media.md`color: red;`
            const query = result[1]

            expect(query).toEqual(
                '@media (min-width: 768px) and (max-width: 1023.95px)',
            )
        })

        it('Should pass through css', () => {
            const query = media.md`color: red;`
            // css`` function is called twice, so we have to dig a bit to fetch the passed value
            const actualCss = query[2][0][0]

            expect(actualCss).toEqual('color: red;')
        })
    })
})
