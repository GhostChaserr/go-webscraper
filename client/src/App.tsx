import { FC } from 'react'
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

import Home from './pages/Home';
import Preview from './pages/Preview';
import { useQuery } from 'react-query';
import axios from 'axios';
import config from './config';
import Cookies from 'js-cookie';

const App: FC = () => {
  const { isLoading, isError, data } = useQuery(`session`, () => axios.get(config.API_URL + `/session`))

  if (isLoading) return (<>Loading</>)
  if (isError) return <div>App is down!</div>

  const newSessionId = data?.data.sessionId as string
  let sessionId = Cookies.get('sessionId')
  if (!sessionId) {
    Cookies.set('sessionId', newSessionId)
    sessionId = newSessionId
  }
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/preview" element={<Preview />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App