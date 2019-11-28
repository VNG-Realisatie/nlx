/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

import React, {useEffect, useState, useRef} from 'react';
import classnames from 'classnames';
import Highlight, {defaultProps} from 'prism-react-renderer';
import defaultTheme from 'prism-react-renderer/themes/palenight';
import Clipboard from 'clipboard';
import rangeParser from 'parse-numeric-range';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './styles.module.css';

import Prism from 'prismjs/components/prism-core';
// For every language that is known to be used in our documentation, load the language definition
import 'prismjs/components/prism-toml.js';
import 'prismjs/components/prism-bash.js';
import 'prismjs/components/prism-powershell.js';
import 'prismjs/components/prism-json.js';

const highlightLinesRangeRegex = /{([\d,-]+)}/;

export default ({children, className: languageClassName, metastring}) => {
  const {
    siteConfig: {
      themeConfig: {prism = {}},
    },
  } = useDocusaurusContext();
  const [showCopied, setShowCopied] = useState(false);
  const target = useRef(null);
  const button = useRef(null);
  let highlightLines = [];

  if (metastring && highlightLinesRangeRegex.test(metastring)) {
    const highlightLinesRange = metastring.match(highlightLinesRangeRegex)[1];
    highlightLines = rangeParser.parse(highlightLinesRange).filter(n => n > 0);
  }

  useEffect(() => {
    let clipboard;

    if (button.current) {
      clipboard = new Clipboard(button.current, {
        target: () => target.current,
      });
    }

    return () => {
      if (clipboard) {
        clipboard.destroy();
      }
    };
  }, [button.current, target.current]);

  let language =
    languageClassName && languageClassName.replace(/language-/, '');

  if (!language && prism.defaultLanguage) {
    language = prism.defaultLanguage;
  }

  const handleCopyCode = () => {
    window.getSelection().empty();
    setShowCopied(true);

    setTimeout(() => setShowCopied(false), 2000);
  };

  return (
    <Highlight
      {...defaultProps}
      Prism={Prism}
      theme={prism.theme || defaultTheme}
      code={children.trim()}
      language={language}>
      {({className, style, tokens, getLineProps, getTokenProps}) => (
        <div className={styles.codeBlockWrapper}>
          <pre
            ref={target}
            className={classnames(className, styles.codeBlock)}
            style={style}>
            {tokens.map((line, i) => {
              const lineProps = getLineProps({line, key: i});

              if (highlightLines.includes(i + 1)) {
                lineProps.className = `${lineProps.className} docusaurus-highlight-code-line`;
              }

              return (
                <div key={i} {...lineProps}>
                  {line.map((token, key) => (
                    <span key={key} {...getTokenProps({token, key})} />
                  ))}
                </div>
              );
            })}
          </pre>
          <button
            ref={button}
            type="button"
            aria-label="Copy code to clipboard"
            className={styles.copyButton}
            onClick={handleCopyCode}>
            <div>
              <svg width="12" height="12" viewBox="340 364 14 15" xmlns="http://www.w3.org/2000/svg">
                <path fill="currentColor"
                      d="M342 375.974h4v.998h-4v-.998zm5-5.987h-5v.998h5v-.998zm2 2.994v-1.995l-3 2.993 3 2.994v-1.996h5v-1.995h-5zm-4.5-.997H342v.998h2.5v-.997zm-2.5 2.993h2.5v-.998H342v.998zm9 .998h1v1.996c-.016.28-.11.514-.297.702-.187.187-.422.28-.703.296h-10c-.547 0-1-.452-1-.998v-10.976c0-.546.453-.998 1-.998h3c0-1.107.89-1.996 2-1.996 1.11 0 2 .89 2 1.996h3c.547 0 1 .452 1 .998v4.99h-1v-2.995h-10v8.98h10v-1.996zm-9-7.983h8c0-.544-.453-.996-1-.996h-1c-.547 0-1-.453-1-.998 0-.546-.453-.998-1-.998-.547 0-1 .452-1 .998 0 .545-.453.998-1 .998h-1c-.547 0-1 .452-1 .997z"
                      fillRule="evenodd"/>
              </svg>
              <strong>{showCopied ? 'Copied' : 'Copy'}</strong>
            </div>
          </button>
        </div>
      )}
    </Highlight>
  );
};
