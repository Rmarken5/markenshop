# Use the official PostgreSQL image as the base image
FROM postgres:15

# Set the working directory
WORKDIR /docker-entrypoint-initdb.d/

# Copy the SQL script to initialize the database
COPY migrations/init.sql .

# Set environment variables for the database
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=dbname
ENV PGDATA /var/lib/webhook/postgresql/data

VOLUME $PGDATA
# Set the default command to run when the container starts
CMD ["postgres"]
