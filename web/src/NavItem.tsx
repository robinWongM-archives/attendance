import React from "react";
import HomeIcon from "./icons/home.svg";

function NavItem() {
  return (
    <div className="text-center text-white cursor-pointer">
      <img src={HomeIcon} className="w-8 h-8 mx-auto fill-current"></img>
      <span className="text-xs">首页</span>
    </div>
  );
}

export default NavItem;
