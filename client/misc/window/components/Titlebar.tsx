/*
 * Not fully impelmented yet, see https://github.com/guasam/electron-window for reference
 */

import React, {
  createRef,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";
import titlebarMenus from "../titlebarMenus";
import context from "../titlebarContextApi";
import { WindowContext } from "./WindowFrame";

type Props = {};

const Titlebar: React.FC<Props> = (props) => {
  const execTest = () => {
    context.exit();
  };

  return (
    <div>
      <button onClick={() => execTest()}>Exit test</button>
    </div>
  );
};

export default Titlebar;
