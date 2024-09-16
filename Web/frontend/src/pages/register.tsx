import { useState } from "react";
import { FaEye, FaEyeSlash } from "react-icons/fa"; 
const Register = () => {
  const [visible , setVisibility] = useState(false)
  
  const handlePasswordVisibility = () => {
    setVisibility(!visible)
  }
  return (
    <div className=" h-screen w-screen flex justify-center items-center">

      <div className=" h-[90%] w-[70%] flex shadow-custom-shadow ">

        <div className="w-[50%] h-full bg-emerald-300 py-10 text-center">
            <h1 className=" text-[36pt]  font-extrabold text-[#234b48]"> Welcome to ShopOps</h1>
            <p className="text-[#234b48] text-[16pt] font-semibold mt-6"> Streamline Your Shop Operations</p>
            <div className="mt-9">
            <p className="text-[12pt] mb-3 text-[#234b48]"> Already have an account?</p>
            <button className="bg-white py-1 px-5 rounded-[10px] text-[#234b48] font-bold text-[12pt]">
                    {" "} Login
                  </button>

            </div>


         
        </div>
        
        <div className="w-[50%] h-full ">
          <div className=" w-[100%] h-full py-4 ">
            <div className="text-center">
              <h1 className="text-[18pt] text-[rgb(90,193,186)] font-bold"> Sign Up</h1>
            </div>

            <div className="px-10">
              <form className="flex flex-col gap-1">
                <div className="flex flex-col ">
                  <label className="text-[10pt] text-gray-600">
                    {" "}
                    First Name
                  </label>
                  <input
                    type="text"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                    
                  />
                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600">
                    {" "}
                    Last Name
                  </label>
                  <input
                    type="text"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />
                        

                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600"> Email</label>
                  <input
                    type="email"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />
                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600">Gender</label>
                  <select className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]">
                  <option value="" selected  hidden> </option>
                    <option>
                      Female
                    </option>
                    <option>
                      Male
                    </option>
                  </select>
              
                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600">
                    Phone Number
                  </label>
                  <input
                    type="tel"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />
                  
                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600"> Password</label>
                  <div className="relative">
                    <input
                      type= {visible ? "text" : "password"} 
                      
                      className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt] w-full pr-10 text-base border border-gray-300 "
                      
                    />
                    <div 
                      className="absolute right-3 top-2 cursor-pointer" onClick={handlePasswordVisibility}

                    >
                      {visible ? <FaEyeSlash /> : <FaEye />}
                      
                    </div>
                  </div>
           
                </div>
                <div className="flex flex-col mb-1">
                  <label className="text-[10pt] text-gray-600">
                    {" "}
                    Confirm Password
                  </label>
                  <input
                    type="text"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />
                </div>

                <div className="flex justify-center  ">
                  <button className="bg-[#009C86]  py-1.5 px-8 rounded-[10px] text-white font-bold text-[12pt]">
                    {" "}
                    Sign Up
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
