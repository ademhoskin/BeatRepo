package repo.beat.middleware.service;

import org.springframework.http.ResponseEntity;

public interface ZipFileS3Service {
    public ResponseEntity<String> uploadZipFile(String bucketName, String objectKey, String filePath);
    public ResponseEntity<String> downloadZipFile(String bucketName, String objectKey, String filePath);
    public ResponseEntity<String> deleteZipFile(String bucketName, String keyName);
}
