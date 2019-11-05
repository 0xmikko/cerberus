import React from 'react';
import {Helmet} from "react-helmet";

export default function() {
  return (
    <div className="App">
      <Helmet>
        <title>No Web3 Provider</title>
      </Helmet>
      <header className="App-header">
        <div>
          No Web3 provider was found. Please, connect Metamask.
          <br />
          More info:{' '}
          <a href={'https://metamask.io/'} style={{color: 'white'}}>
            metamask.io
          </a>
        </div>
      </header>
    </div>
  );
}
