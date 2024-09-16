import { useState } from "react";
import { FaEye, FaEyeSlash } from "react-icons/fa"; 
import {Link} from 'react-router-dom'
import {useForm , SubmitHandler} from 'react-hook-form'
import {z} from 'zod';
import {zodResolver} from "@hookform/resolvers/zod";


const Register = () => {
  const schema = z.object({
    firstName: z.string().min(1 , "first name is required"),
    lastName:z.string().min(1 , "last name is required"),
    email:z.string().email(),
    gender: z.enum(["Male", "Female"], {
      invalid_type_error: "Select a valid gender", 
    }),
    phoneNumber: z.string().regex( /(\+\s*2\s*5\s*1\s*9\s*(([0-9]\s*){8}\s*))|(0\s*9\s*(([0-9]\s*){8}))/ , "Invalid phone number format"),
    password:z.string().min(8).regex(/[A-Z]/, 'Password must contain at least one uppercase letter')
    .regex(/[!@#$%^&*(),.?":{}|<>]/, 'Password must contain at least one special character'),
    confirmPassword: z.string(),

  }).refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    message: "Passwords don't match",
  });

 type fields = z.infer<typeof schema>
 const { register , handleSubmit , formState:{errors}, clearErrors} = useForm<fields>(
  {resolver: zodResolver(schema)}
)
  const [visible , setVisibility] = useState(false)
  const handlePasswordVisibility = () => {
    setVisibility(!visible)
  }
  const onSubmit: SubmitHandler<fields> = (data) =>{
    console.log(data)
  }

 

  return (
    <div className=" h-screen w-screen flex justify-center items-center">

      <div className=" h-[90%] w-[70%] flex shadow-custom-shadow ">

        <div className="w-[50%] h-full bg-emerald-300 py-10 text-center">
            <h1 className=" text-[36pt]  font-extrabold text-[#234b48]"> Welcome to ShopOps</h1>
            <p className="text-[#234b48] text-[16pt] font-semibold mt-6"> Streamline Your Shop Operations</p>
            <div className="mt-9">
            <p className="text-[12pt] mb-3 text-[#234b48]"> Already have an account?</p>
            <Link to = "/login" className="bg-white py-1 px-5 rounded-[10px] text-[#234b48] font-bold text-[12pt]">
                    {" "} Login
                  </Link>

            </div>


         
        </div>
        
        <div className="w-[50%] h-full ">
          <div className=" w-[100%] h-full py-4 ">
            <div className="text-center">
              <h1 className="text-[18pt] text-[rgb(90,193,186)] font-bold"> Sign Up</h1>
            </div>

            <div className="px-10">
              <form className="flex flex-col gap-1" onSubmit={handleSubmit(onSubmit)}>
                <div className="flex flex-col ">
                <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600">First Name</label>
                  {errors.firstName && <p className="text-[10pt] text-red-600">{errors.firstName.message}</p>}
                  </div>
                  <input
                  {...register('firstName')}
                    type="text"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                    
                  />
           
                </div>
                <div className="flex flex-col mb-1">
                <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600">Last Name</label>
                  {errors.lastName && <p className="text-[10pt] text-red-600">{errors.lastName.message}</p>}
                  </div>
                  <input
                  {...register('lastName')}
                    type="text"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />
 

                </div>
                <div className="flex flex-col mb-1">
                <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600">First Name</label>
                  {errors.email && <p className="text-[10pt] text-red-600">{errors.email.message}</p>}
                  </div>                  <input
                  {...register('email')}
                    type="email"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />

                </div>
                <div className="flex flex-col mb-1">
                  <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600">Gender</label>
                  {errors.gender && <p className="text-[10pt] text-red-600">select one</p>}
                  </div>
                
                  <select className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  {...register('gender')}>
                  <option value=""   hidden> </option>
                    <option>
                      Female
                    </option>
                    <option>
                      Male
                    </option>
                  </select>
                  

              
                </div>
                <div className="flex flex-col mb-1">
                <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600"> Phone Number</label>
                  {errors.phoneNumber && <p className="text-[10pt] text-red-600">{errors.phoneNumber.message}</p>}
                  </div>
                  <input
                    type="tel"
                    {...register('phoneNumber')}
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />

                </div>
                <div className="flex flex-col mb-1">
                <div className="flex justify-between gap-5">
                  <label className="text-[10pt] text-gray-600">Password</label>
                  {errors.password && <p className="text-[10pt] text-red-600">{errors.password.message}</p>}
                  </div>                  <div className="relative">
                    <input
                    {...register('password')}
                      type= {visible ? "text" : "password"} 
                      
                      className="h-8 bg-[#F1F5F5] border-none rounded-[5px] outline-none px-3 text-[11pt] w-full pr-10 text-base border border-gray-300 "
                      
                    />
                    <div 
                      className="absolute right-3 top-2 cursor-pointer" onClick={handlePasswordVisibility}

                    >
                      {visible ? <FaEyeSlash /> : <FaEye />}
                      
                    </div>

                  </div>
           
                </div>
                <div className="flex flex-col mb-1">
                <div className="flex justify-between">
                  <label className="text-[10pt] text-gray-600">Confirm Password</label>
                  {errors.confirmPassword && <p className="text-[10pt] text-red-600">{errors.confirmPassword.message}</p>}
                  </div>
                  <input
                  {...register('confirmPassword')}
                    type="password"
                    className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[11pt]"
                  />

                </div>

                <div className="flex justify-center  ">
                  <button type="submit" className="bg-[#009C86]  py-1.5 px-8 rounded-[10px] text-white font-bold text-[12pt]">
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
