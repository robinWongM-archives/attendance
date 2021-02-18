import React from "react";
import Logo from "./ecnc.svg";
import NavItem from "./NavItem";
import ProfileIcon from "./icons/profile.svg";

function Nav() {
  return (
    <div className="h-screen w-16 py-6 flex flex-col items-center flex-shrink-0">
      <img src={Logo} className="w-12 h-12"></img>
      <div className="flex-1 flex flex-col justify-center">
        <NavItem></NavItem>
      </div>
      <img src={ProfileIcon} className="w-10 h-10"></img>
    </div>
  );
}

export default Nav;
