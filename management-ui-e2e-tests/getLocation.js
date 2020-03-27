import { ClientFunction } from 'testcafe';

const getLocation = ClientFunction(() => document.location.href);

module.exports = getLocation
