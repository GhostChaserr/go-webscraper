import { FC, ReactNode } from 'react'

type FieldCardProps = {
  available: boolean
  field: string
  content: ReactNode
}

const FieldCard:FC<FieldCardProps> = ({ available, field, content }) => {
  return (
    <div>
      Field Card
    </div>
  )
}

export default FieldCard