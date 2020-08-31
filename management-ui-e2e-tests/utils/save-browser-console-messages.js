// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { t } from 'testcafe'
import fs from 'fs'
import path from 'path'

export default async function ({ testController, directory, fileName }) {
  const browserConsoleMessages = await t.getBrowserConsoleMessages()
  const filePath = path.join(directory, fileName)

  const result = {
    fixture: testController.testRun.test.testFile.currentFixture.name,
    testName: testController.testRun.test.name,
    browserConsoleMessages: browserConsoleMessages
  }

  await fs.promises.writeFile(filePath, JSON.stringify(result, null, 4))
}
