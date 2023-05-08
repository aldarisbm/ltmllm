# LTMLLM

ltmllm is a simple backend to create a long term memory large language model

### Disclaimer
This is a work in progress and is not ready for use. I will only support
OpenAI GPT streaming models.

## Details:
The way this work is as follows:
- We create a llm chatbot. 
- We will create a local k/v database(bolt) to save full messages and conversations
- We will use an embedding API to create embeddings for each prompt and response
- We will store the embeddings in a vector db (pinecone) and metadata pointing to bolt
- Every time we interact with our chatbot, it will embed our question
and query the vector db for the most similar messages
- We will then retrieve the full message to give the chatbot context


