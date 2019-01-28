import { relativeToFullUrl } from './api'

describe('api helper', () => {
    describe('convert relative to full url', () => {
        describe('without REACT_APP_API_BASE_URL being set', () => {
            it('should prefix with /api by default', () => {
                delete process.env.REACT_APP_API_BASE_URL
                expect(relativeToFullUrl('/out')).toBe('/api/out')
            })
        })

        describe('with REACT_APP_API_BASE_URL defined', () => {
            it('should prefix with /api if the REACT_APP_API_BASE_URL variable is not set', () => {
                process.env.REACT_APP_API_BASE_URL = 'https://www.duck.com'
                expect(relativeToFullUrl('/out')).toBe(
                    'https://www.duck.com/out',
                )
            })
        })
    })
})
