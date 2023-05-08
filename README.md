# LTMLLM

ltmllm is a simple backend to create a long term memory large language model

### Disclaimer
This is a work in progress and is not ready for use. I will only support
then OpenAI GPT Chat stream models, and Pinecone embeddings API. I will use a 
boltdb database to store the full messages and conversations. I will use a 
pinecone vector database to store the embeddings.

## Details:
The way this work is as follows:
- We have an llm chatbot created
- We will create a local database to save full messages and conversations
- We will use an embedding API to create embeddings for each message
- We will store the embeddings in a vector db (pinecone)
- Every time we interact with our chatbot, it will embed our question
and query the vector db for the most similar messages
- We will then query our local database for the full messages
- We will then use the full messages to give the chatbot context


