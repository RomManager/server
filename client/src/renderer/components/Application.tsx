import React from "react";
import logo from "@assets/icons/electron-logo.png";

const Application: React.FC = () => {
  return (
    <div>
      <h1>Welcome to the RM Client</h1>
      <img src={logo} style={{ width: 100 }} />
    </div>
  );
};

export default Application;
