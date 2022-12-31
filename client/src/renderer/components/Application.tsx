import React from 'react';
import logo from '@assets/icons/chrome.png'

const Application: React.FC = () => {
  return (
    <div>
      <h1>Welcome to the RM Client</h1>
      <img src={logo}/>
    </div>
  );
};

export default Application;
