import React from "react";
import HomeIcon from "./icons/home.svg";

function NavItem() {
  return (
    <div className="text-center text-white">
      <img src={HomeIcon} className="w-10 h-10 mx-auto fill-current"></img>
      <span className="text-sm">首页</span>
    </div>
  );
}

export default NavItem;
