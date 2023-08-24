# TODO

## To Validate

Validate companies are interested in labelling/verification done by experts.

  - Data labelling vs verifying other AI models
  - Time expectation? Willing to pay surge to get it done faster?
  - Is it important to be able to send same job to more than one expert to verify is correct?
  - What expertise do they want? Coder, medical, language...
  - Integration into their systems important?
  - Do they want tool to validate/label their own data with their labellers?
  - Do they want experts with high confidence of expertise (while paying more) or are ok with CV check while paying less?
  - Amount of labelling?
  - Training open source models with labelled data? or just raw labelling?
  - How often they need labelling job?
  - RLHF, 
  - ML model evaluation

## Features

- [ ] Create labelling project [Companies]
  - Name
  - Status
  - Company
  - private bool (if private, only people with link can respond)
  - instructions
  - workers per task
  - qualifications_required
  - callback_url
  - tags
  - [question type](https://app.surgehq.ai/docs/api)
    - multiple choice - select 1
      - text
      - options
      - descriptions
      - require_tiebreaker
      - column_header
      - preexisting_annotations
    - multiple choice - select many
    - free response
    - text tagging
      - round nearest word
    - Tree Selection
    - Text Area
    - File Upload
    - Ranking
    - ChatBot
- List projects
- Retrieve a project
- Set status project
- Sample responses

- [x] some




Entities

- Project
- Question
    - Multiple Choice
    - Free Response
    - Text Tagging
    - Chatbot
    - Ranking
- Response
- Expert
- Qualification

TODO:

- Project
    - Create, Read 1, List, Update
- Question
    - Multiple Choice
    - Single Choice
    - Free Response
    - Text Tagging
    - Chatbot
    - Ranking
- Expert
    - Create, Update
- Response
    - Create
