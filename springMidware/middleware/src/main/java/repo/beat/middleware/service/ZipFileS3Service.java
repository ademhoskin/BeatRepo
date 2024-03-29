package repo.beat.middleware.service;

import org.springframework.http.ResponseEntity;

public interface ZipFileS3Service {
     ResponseEntity<String> uploadZipFile(String bucketName, String objectKey, String filePath);
     ResponseEntity<String> downloadZipFile(String bucketName, String objectKey, String filePath);
     ResponseEntity<String> deleteZipFile(String bucketName, String keyName);
}
