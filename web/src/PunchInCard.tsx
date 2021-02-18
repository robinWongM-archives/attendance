import { Button } from "./Button";
import React from "react";
import ApproveIcon from "./icons/approve.svg";
import CancelIcon from "./icons/cancel.svg";

function PunchInCard() {
  return (
    <div className="w-max bg-gradient-to-r from-blue-500 to-blue-300 rounded-2xl mt-8 shadow-xl px-8 py-6 text-white flex">
      <div>
        <p className="text-opacity-90 text-base">
          点击签到前，请确认以下信息是否正确。
        </p>
        <div className="mt-8 mr-8 flex">
          <div className="mr-8">
            <div className="text-white text-opacity-60 text-base">日期</div>
            <div className="text-white text-opacity-100 text-2xl">
              12 月 28 日
            </div>
          </div>
          <div className="mr-8">
            <div className="text-white text-opacity-60 text-base">岗位</div>
            <div className="text-white text-opacity-100 text-2xl">电话岗</div>
          </div>
          <div className="mr-8">
            <div className="text-white text-opacity-60 text-base">时间段</div>
            <div className="text-white text-opacity-100 text-2xl">
              09:00 ~ 10:30
            </div>
          </div>
          <div className="mr-8">
            <div className="text-white text-opacity-60 text-base">代班</div>
            <div className="text-white text-opacity-100 text-2xl">无</div>
          </div>
        </div>
      </div>
      <div className="float-left flex flex-col">
        <Button className="mb-5" Icon={ApproveIcon} Content="确认签到" />
        <Button Icon={CancelIcon} Content="信息有误" />
      </div>
    </div>
  );
}

export default PunchInCard;
