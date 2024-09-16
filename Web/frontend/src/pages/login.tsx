import { useState } from "react";
import person from "../assets/person.png";
import { Link } from "react-router-dom";
import { FaEye, FaEyeSlash } from "react-icons/fa";
import {useForm , SubmitHandler} from 'react-hook-form'
import {z} from 'zod';
import {zodResolver} from "@hookform/resolvers/zod";


const Login = () => {
  const schema = z.object({
    email:z.string().email(),
    phoneNumber: z.string().regex( /(\+\s*2\s*5\s*1\s*9\s*(([0-9]\s*){8}\s*))|(0\s*9\s*(([0-9]\s*){8}))/ , "Invalid phone number format"),
    password:z.string().min(8).regex(/[A-Z]/, 'Password must contain at least one uppercase letter')
    .regex(/[!@#$%^&*(),.?":{}|<>]/, 'Password must contain at least one special character'),
  })

  type fields = z.infer<typeof schema>

  const { register , handleSubmit , formState:{errors}, clearErrors} = useForm<fields>(
    {resolver: zodResolver(schema)}
  )

  const onSubmit: SubmitHandler<fields> = (data) =>{
    console.log(data)
  }

  const [visible , setVisibility] = useState(false)
  const handlePasswordVisibility = () => {
    setVisibility(!visible)
  }

  return (
    <div className=" w-screen h-screen  flex  gap-6">
      <div className=" w-[40%] h-full m-11">
        <img src={person} />
      </div>
      <div className=" w-[60%] h-full flex py-1">
        <div className="w-[60%] h-[80%] shadow-custom-shadow p-2 my-10  flex flex-col ">
          <div className="flex flex-col items-center">
            <h1 className="text-[25pt]  font-bold">
              {" "}
              Shop<span className="text-[#00C8AC]">Ops</span>
            </h1>
            <p className="text-[#234b48] text-[11pt] mt-3 ">
              {" "}
              Streamline Your Shop Operations
            </p>
            <h1 className="text-[18pt] text-[#5ac1ba] font-bold mt-2">Login</h1>
          </div>

          <div className="px-7 mt-3">
            <form onSubmit={handleSubmit(onSubmit)}>
              <div className="flex flex-col mb-1">
                <label className="text-[10pt] text-gray-600"> Email</label>
                <input
                {...register("email")}
                type="email"
                // onChange={() => clearErrors('email')}
                  className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[10pt]"
                />
                {errors.email &&<p className="text-red-600 text-[10pt]"> {errors.email.message}</p> }
              </div>

              <div className="flex flex-col mb-1">
                <label className="text-[10pt] text-gray-600">
                  {" "}
                  Phone Number
                </label>
                <input
                  type="tel"
                   {...register("phoneNumber")}
                  className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[10pt]"
                />
                 {errors.phoneNumber &&<p className="text-red-600 text-[10pt]"> {errors.phoneNumber.message}</p> }
              </div>
              <div className="flex flex-col mb-1">
              <label className="text-[10pt] text-gray-600"> Password</label>

              <div className="relative">
                    <input
                      type= {visible ? "text" : "password"} 
                      {...register("password" )}
                      
                      className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none border-none px-3 text-[11pt] w-full pr-10 text-base border border-gray-300 "
                      
                    />
                    <div 
                      className="absolute right-3 top-2 cursor-pointer" onClick={handlePasswordVisibility}

                    >
                      {visible ? <FaEyeSlash /> : <FaEye />}
                      
                    </div>
                  </div>
                  {errors.password &&<p className="text-red-600 text-[10pt]"> {errors.password.message}</p> }
                  </div>
              <div className="flex justify-start ">
                <p className="text-[10pt] mt-2 mb-5">
                  {" "}
                  Don't have an account ?{" "}
                  <Link to="/register" className="text-[#5ac1ba] font-bold">
                    {" "}
                    Sign Up Here
                  </Link>{" "}
                </p>
              </div>
              <div className="flex justify-center">
                <button className="bg-[#009C86] pb-1 px-8 rounded-[10px] text-white font-bold text-[13pt] text-center mb-2" >
                  Login
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
