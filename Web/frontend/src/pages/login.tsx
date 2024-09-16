import person from '../assets/person.png'
const Login = () => {
  return (
    <div className=" w-screen h-screen  flex  gap-6"> 
        <div className=" w-[40%] h-full m-11">
            <img src={person} />

        </div>
        <div className=" w-[60%] h-full flex py-1">
            <div className="w-[60%] h-[80%] shadow-custom-shadow p-2 my-10  flex flex-col ">
                <div className="flex flex-col items-center">
                <h1 className="text-[25pt]  font-bold" > Shop<span className="text-[#00C8AC]">Ops</span></h1>
                <p className="text-[#234b48] text-[11pt] mt-3 "> Streamline Your Shop Operations</p>
                <h1 className="text-[18pt] text-[#5ac1ba] font-bold mt-2">Login</h1>
                </div>

            <div className="px-7 mt-9">
                <form>
               
                  
                    <div className="flex flex-col mb-1">
                    <label className="text-[10pt] text-gray-600"> Email</label>
                    <input type='email' className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[10pt]" />
                    </div>
               
                    <div className="flex flex-col mb-2">
                    <label className="text-[10pt] text-gray-600"> Password</label>
                    <input type="text" className="h-8 bg-[#F1F5F5] rounded-[5px] outline-none px-3 text-[10pt]" />
                    </div>
                    <div className="flex justify-start ">
                    <p className="text-[11pt]"> Already have an account? <span className='text-[#5ac1ba] font-bold'> Login Here</span> </p>
                    </div>
                     <div className="flex justify-center mt-9 ">
                     <button className="bg-[#009C86] pb-1 px-8 rounded-[10px] text-white font-bold text-[13pt] text-center">Login</button>
                     </div>
                    
                  
                    
                </form>
            </div>
            </div>


        </div>
    </div>
  )
}

export default Login