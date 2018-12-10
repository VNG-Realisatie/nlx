import build from '../build'

import parseColorDerivates from '../parseColorDerivates'

jest.mock('../parseColorDerivates')
jest.mock('../themeConstants')
jest.mock('../themeDefault')

describe('themes/build', () => {
    it('should remember the generated style object', () => {
        build()
        build()
        expect(parseColorDerivates).toHaveBeenCalledTimes(1)
    })
})
