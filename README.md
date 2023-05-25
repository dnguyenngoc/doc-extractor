# Document Processing Pipeline

This project demonstrates how to build a document processing streaming pipeline. The base pipeline provided here allows for easy implementation of additional features using machine learning models as workers in the pipeline.

We have chosen invoice processing as the base pipeline and this release serves as the initial codebase to handle it. The base pipeline includes the following components:

- **API**: Provides an interface for customers to upload documents with supported types (JPEG, PNG, TIF).
- **Image Processing**: Preprocesses the invoice image to enhance quality, remove noise, and improve readability if necessary. Techniques such as image cropping, rotation, or resizing are applied to isolate and align the invoice content.
- **Field Invoice Detection**: Utilizes techniques like object detection or contour analysis to identify and extract the invoice region within the processed image. This step aims to isolate the invoice from any surrounding elements or backgrounds.
- **OCR (Optical Character Recognition)**: Applies OCR techniques to recognize and extract text from the invoice image or the specific invoice region identified in the previous step. We have used an OCR model built using machine learning.
- **Information Extraction and Parsing**: Performs calculations or data processing on the extracted information if necessary. For example, calculating the total amount, taxes, or applying business-specific rules.

## Technologies Used

The following technologies are used in this project:

1. **Golang** with the **gorillamux** framework for building the API and handling public endpoints.
2. **S3**: Used for storage and handling of image and JSON files.
3. **Postgrest**: A database used for storing user information, document processing details, and other relevant data.
4. **Kafka**: A messaging service used for asynchronous communication between different components of the pipeline.
5. **Redis**: A caching service used to improve performance and efficiency.
6. **Faust**: A framework for streaming processing, which helps in building efficient and scalable pipeline components.

## Getting Started

To get started with the document processing pipeline, follow these steps:

1. Clone the repository.
2. Install the required dependencies and ensure all necessary services (S3, Postgrest, Kafka, Redis) are set up and running.
3. Configure the project settings, including API keys, database connections, and messaging service details.
4. Build and run the project using the provided commands.
5. Access the API endpoints to upload documents and initiate the document processing pipeline.

## Example Pipeline Process

![Pipeline Process](/document/images/pipeline_sample.png)

## Contact

For any inquiries or assistance, please contact our team at [email protected] or visit our website at www.example.com.

We appreciate your interest in our document processing pipeline!
