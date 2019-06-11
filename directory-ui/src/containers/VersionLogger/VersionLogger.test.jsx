import React from 'react';
import {shallow, mount} from 'enzyme';
import VersionLogger from './VersionLogger';
import { act, create } from "react-test-renderer";

describe('VersionLogger', () => {
    beforeEach(() => {
        fetch.resetMocks()
        fetch.mockResponseOnce(JSON.stringify({ tag: 'test' }))
    })

    describe('effect', () => {
        it('should call logger when rendered', (done) => {
            const logger = jest.fn()
            act(() => {
                // the async call in useEffect will cause a warning. this will pro
                create(<VersionLogger logger={logger}/>)
            })
            expect(fetch).toHaveBeenCalledTimes(1)
            
            // workaround for async result of fetch
            setTimeout(() => {
                expect(logger).toHaveBeenCalledWith('test')
                done()
            })
        })
    })

    describe('snapshot', () => {
        it('should should match snapshot', () => {
            const logger = jest.fn()

            expect(shallow(<VersionLogger logger={logger}/>)).toMatchSnapshot()
        });
    });
});