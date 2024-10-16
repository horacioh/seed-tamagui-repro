import {TamaguiProvider, config} from '@shm/ui'
import '@tamagui/core/reset.css'
import ReactDOM from 'react-dom/client'
import {App} from './app'
import './root.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <TamaguiProvider config={config}>
    <App />
  </TamaguiProvider>,
)
